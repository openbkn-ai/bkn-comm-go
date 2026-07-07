package oteltrace

import (
	"context"
	"fmt"
	"net/http"
	"runtime"
	"strings"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/otel"
	attr "go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
)

const (
	// InstrumentationName 用于创建 tracer 的 instrumentation name
	InstrumentationName = "bkn-backend/otel"
)

// StartInternalSpan 服务内函数调用创建 span，自动从 runtime.Caller 获取 span name。
func StartInternalSpan(ctx context.Context) (context.Context, trace.Span) {
	name, filepath := callerFuncName(2)
	newCtx, span := StartNamedInternalSpan(ctx, name)
	if filepath != "" {
		span.SetAttributes(attr.String("code.filepath", filepath))
	}
	return newCtx, span
}

// StartClientSpan 外部依赖调用创建 span，自动从 runtime.Caller 获取 span name。
func StartClientSpan(ctx context.Context) (context.Context, trace.Span) {
	name, filepath := callerFuncName(2)
	newCtx, span := StartNamedClientSpan(ctx, name)
	if filepath != "" {
		span.SetAttributes(attr.String("code.filepath", filepath))
	}
	return newCtx, span
}

func callerFuncName(skip int) (string, string) {
	pc, file, lineNo, ok := runtime.Caller(skip)
	if !ok {
		return "unknown", ""
	}
	funcPaths := strings.Split(runtime.FuncForPC(pc).Name(), "/")
	return funcPaths[len(funcPaths)-1], fmt.Sprintf("%s:%v", file, lineNo)
}

// StartNamedClientSpan用自定义业务名创建 SpanKindClient 类型 span。
func StartNamedClientSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	return otel.Tracer(InstrumentationName).Start(ctx, name, trace.WithSpanKind(trace.SpanKindClient))
}

// StartNamedInternalSpan 用自定义业务名创建 SpanKindInternal 类型 span。
func StartNamedInternalSpan(ctx context.Context, name string) (context.Context, trace.Span) {
	return otel.Tracer(InstrumentationName).Start(ctx, name, trace.WithSpanKind(trace.SpanKindInternal))
}

// StartServerSpan 跨服务（HTTP 接口）创建 span。
// 前置依赖：TracingMiddleware 已把 trace header 提取到 c.Request.Context()，
//
//	LanguageMiddleware 已把 language 叠加到 c.Request.Context()。
func StartServerSpan(c *gin.Context) (context.Context, trace.Span) {
	spanName := fmt.Sprintf("%s %s", c.Request.Method, c.FullPath())
	newCtx, span := otel.Tracer(InstrumentationName).Start(c.Request.Context(), spanName, trace.WithSpanKind(trace.SpanKindServer))
	span.SetAttributes(
		attr.String("http.request.method", c.Request.Method),
		attr.String("http.route", c.FullPath()),
		attr.String("client.address", c.ClientIP()),
	)

	return newCtx, span
}

// ExtractTraceHeader 从 HTTP Header 中提取 Trace 上下文。
func ExtractTraceHeader(ctx context.Context, header http.Header) context.Context {
	if header == nil {
		return ctx
	}

	return otel.GetTextMapPropagator().Extract(ctx, propagation.HeaderCarrier(header))
}

// SetAttributes 在当前 span 上设置属性。
func SetAttributes(ctx context.Context, kv ...attr.KeyValue) {
	span := trace.SpanFromContext(ctx)
	span.SetAttributes(kv...)
}

// EndSpan 结束当前 span，如有错误则记录。
func EndSpan(ctx context.Context, err error) {
	span := trace.SpanFromContext(ctx)
	if span == nil {
		return
	}

	if err != nil {
		span.RecordError(err)
		span.SetStatus(codes.Error, err.Error())
	} else {
		span.SetStatus(codes.Ok, "OK")
	}

	span.End()
}
