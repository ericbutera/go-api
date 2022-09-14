// https://medium.com/jaegertracing/take-jaeger-for-a-hotrod-ride-233cf43e46c2
// https://github.com/jaegertracing/jaeger/tree/main/examples/hotrod

package api

import (
	"log"

	// appconfig "github.com/ericbutera/go-api/config"
	// "github.com/ericbutera/go-api/api"

	"github.com/gin-gonic/gin"
	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/jaeger"
	"go.opentelemetry.io/otel/sdk/resource"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func InitOpenTel(app *App, r *gin.Engine) {
	tp, err := OpenTelProvider(app)
	if err != nil {
		log.Fatal(err)
	}

	// TODO: something about this doesn't work:
	// ctx, cancel := context.WithCancel(context.Background())
	// defer cancel()
	// // Cleanly shutdown and flush telemetry when the application exits.
	// defer func(ctx context.Context) {
	// 	// Do not make the application hang when it is shutdown.
	// 	ctx, cancel = context.WithTimeout(ctx, time.Second*5)
	// 	defer cancel()
	// 	if err := tp.Shutdown(ctx); err != nil {
	// 		log.Fatal(err)
	// 	}
	// }(ctx)

	//tr := tp.Tracer(config.ServiceName)
	// ctx, span := tr.Start(ctx, "foo")
	// defer span.End()
	r.Use(otelgin.Middleware(app.Config.AppName, otelgin.WithTracerProvider(tp)))
}

// tracerProvider returns an OpenTelemetry TracerProvider configured to use
// the Jaeger exporter that will send spans to the provided url. The returned
// TracerProvider will also use a Resource configured with all the information
// about the application.
func OpenTelProvider(app *App) (*sdktrace.TracerProvider, error) {
	endpoint := jaeger.WithEndpoint(app.Config.JaegerCollectorEndpoint)
	exporter, err := jaeger.New(jaeger.WithCollectorEndpoint(endpoint))
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		// Always be sure to batch in production.
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
		// Record information about this application in a Resource.
		sdktrace.WithResource(resource.NewWithAttributes(
			semconv.SchemaURL,
			semconv.ServiceNameKey.String(app.Config.AppName),
			attribute.String("environment", app.Config.Env),
		)),
	)

	// Register our TracerProvider as the global so any imported
	// instrumentation in the future will default to using it.
	otel.SetTracerProvider(tp)
	// otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))

	return tp, nil
}
