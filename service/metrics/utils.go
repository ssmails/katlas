package metrics

import (
        "github.com/prometheus/client_golang/prometheus"
        dto "github.com/prometheus/client_model/go"
)

func ReadCounter(m prometheus.Counter) float64 {
    pb := &dto.Metric{}
    m.Write(pb)
    return pb.GetCounter().GetValue()
}
