FROM envoyproxy/envoy-dev:0458877b46f9bcb96f843331e2515f7aaf6bb800
COPY envoy.yaml /etc/envoy/envoy.yaml
RUN chmod go+r /etc/envoy/envoy.yaml