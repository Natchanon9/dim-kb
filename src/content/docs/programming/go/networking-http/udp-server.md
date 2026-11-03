---
title: "UDP Server"
description: "Datagrams with net.ListenPacket"
sidebar:
  order: 9
status: done
lastReviewed: 2026-11-03
---

UDP sends independent datagrams with no connection, ordering or delivery guarantee. It suits DNS, metrics (StatsD), games and discovery.

## Example

```go title="udp.go"
pc, err := net.ListenPacket("udp", ":9001")
if err != nil {
	log.Fatal(err)
}
buf := make([]byte, 1500)
for {
	n, addr, err := pc.ReadFrom(buf) // one datagram and its sender
	if err != nil {
		continue
	}
	pc.WriteTo(bytes.ToUpper(buf[:n]), addr)
}
// client sends "ping" → receives "PING"
```

| | TCP | UDP |
| --- | --- | --- |
| Connection | Yes | No |
| Order / delivery | Guaranteed | Not guaranteed |
| Unit | Byte stream | Datagram |

:::caution[Pitfalls]
- A datagram bigger than your buffer is truncated; keep messages under ~1400 bytes to avoid fragmentation.
:::
