# local_llm_cluster
Implementation and configuration of local LLM using Ollama and Cluster of 4 PCs [ and maybe scaling to 16 PCs and 32 PCs ].
For LLM Distributed inference over the network.

## COMPONENTS:
- [x] 4 Computers ai_machine(HEAD) & ai_node1 to ai_node(n):
(3 PC and a supermicro server) and each has:
  - 1 Mellanox MCX354A-FCBT CX354A ConnectX-3 VPI
  - 1 Nvidia Quadro P6000 with 24GB VRAM
  - 128 GB Memory
  - 1 Intel Xeon E5-2690 v4 Tetradeca-core (14 Core), Socket R3 (LGA2011-3)
        Note that the server has a dual socket, Hence, 2X Intel Xeon E5-2690 v4 Tetradeca-core (28 Core)
- [x] Networking:
  - 1 Cisco Switch: CISCO NEXUS 3172PQ with 40Gbps & 56 Gbps
  - 1 Cisco Router: ISR 1100 Series
  - 1 Internal network Router (as Local Enterprise ISP): Asus RT-AX1800S V2
  - 1 Main router: Cox Panoramic(Internet Provider)
  - 8 QSFP40GPC05M MSA Compliant QSFP+ Direct-Attach Twinax Cable - 0.5 m (1.6 ft) - 40 GbE
- [x] AMP up Powerlines:
  - [x] Install Subpanel in Garage
  - [x] Run 20 AWG wires from Subpannel to Home Office
  - [x] Powerup all components: 4 outlets with 20 AMP each
     
## CONGIGURATION:
# Software implementation & configuration
- [x] Install Ubuntu 22.04 on each machine
- [x] Downloand and configure Drivers for the Mellanox ConnectX-3 cards
- [ ] Establish network connection between all PCs at 40Gbps & enable routing for internet connection  
- [ ] Ollama Setup on all machines
- [x] Launch Ollama on one machine
- [x] Test Run LLMs of 26B parameters and less on one PC
- [ ] Launch Ollama on all machines
- [ ] Test Run LLms on all machines seperately
- [ ] Launch Ray Across All Nodes
- [ ] Start vLLM with Tensor Parallelism
- [ ] Run LLM distributed Inference over the network on all machine **ACTING** as one PC 


# Hardware implementation & configuration
- Internal Router ASUS RT-AX1800S configuration:
  - [x] Change default getaway from 192.168.x.x to enterprise IP 10.xx.xx.xx
  - [x] Enable DHCP and VPN connection

- ISR Router configuration:
  - [x] Update software and licenses
  - [ ] Configure out port for QSFP 40Gbps connection  & internet connection (routing) to Cisco Switch


- CISCO NEXUS 3172PQ configuration:
  - [x] Update software and licenses
  - [ ] Configure 4 QSFP ports & Routing for the computers

💠 Working progress ...

