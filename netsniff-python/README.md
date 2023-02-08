<!-- @format -->

# Network Sniffer using Scapy 🔎

This code captures and analyzes network traffic using the Scapy library in Python.

## Code Explanation 💻

The code uses the sniff function from Scapy to capture network packets and the handle_packet function as a callback to process each packet received.

The handle_packet function checks the protocol of each packet using the IP and TCP or UDP headers. For TCP packets, it extracts the source and destination IP addresses and port numbers, as well as the TCP flags, sequence number, and acknowledgment number. For UDP packets, it extracts the source and destination IP addresses and port numbers, as well as the UDP length.

The code also implements filtering of packets based on specific criteria. For example, in the case of TCP packets, it only prints the packet information if the TCP flags are equal to 2. In the case of UDP packets, it only prints the packet information if the destination port number is equal to 53.

## Output Explanation 📈

The output of this code is a list of packets captured from the network, with the following information for each packet:

- Protocol (TCP or UDP)
- Source IP address and port number
- Destination IP address and port number
- Additional information such as TCP flags, sequence number, acknowledgment number, or UDP length

The output is printed to the console, allowing you to view and analyze network traffic in real-time.

## Conclusion 🎉

This code provides a simple yet powerful example of how the Scapy library can be used to capture and analyze network traffic in Python. Whether you're a network engineer, security analyst, or software developer, this code can be a useful starting point for your next network analysis project.
