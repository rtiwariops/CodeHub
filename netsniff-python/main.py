from scapy.all import *

def handle_packet(pkt):
    # Check the protocol of the packet
    if IP in pkt:
        ip_src = pkt[IP].src
        ip_dst = pkt[IP].dst
        if TCP in pkt:
            tcp_sport = pkt[TCP].sport
            tcp_dport = pkt[TCP].dport
            # Analyze other layers of the network stack (e.g., TCP headers)
            tcp_flags = pkt[TCP].flags
            tcp_seq = pkt[TCP].seq
            tcp_ack = pkt[TCP].ack
            # Filter packets based on specific criteria
            if tcp_flags == 2:
                print("Protocol: TCP\nSource IP: {}:{}\nDestination IP: {}:{}\nTCP Flags: {}\nTCP Seq: {}\nTCP Ack: {}".format(ip_src, tcp_sport, ip_dst, tcp_dport, tcp_flags, tcp_seq, tcp_ack))
        elif UDP in pkt:
            udp_sport = pkt[UDP].sport
            udp_dport = pkt[UDP].dport
            # Analyze other layers of the network stack (e.g., UDP headers)
            udp_len = pkt[UDP].len
            # Filter packets based on specific criteria
            if udp_dport == 53:
                print("Protocol: UDP\nSource IP: {}:{}\nDestination IP: {}:{}\nUDP Length: {}".format(ip_src, udp_sport, ip_dst, udp_dport, udp_len))

# Start receiving packets using Scapy's sniff function
sniff(prn=handle_packet)
