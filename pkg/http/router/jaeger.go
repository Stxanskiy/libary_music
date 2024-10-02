package router

/*func SpanInject(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var (
			span   opentracing.Span
			tracer = opentracing.GlobalTracer()
		)
		sc, err := tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(r.Header))
		if err != nil && !errors.Is(err, opentracing.ErrSpanContextNotFound) {
			lg.Warnf("Warning with extracting: %v", err)
		}
		span = tracer.StartSpan(r.URL.Path, opentracing.ChildOf(sc))
		defer span.Finish()
		ext.MessageBusDestination.Set(span, r.URL.Path)
		next.ServeHTTP(w, r.WithContext(opentracing.ContextWithSpan(r.Context(), span)))
	})
}
*/
