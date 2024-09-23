Round Robin algorithm for load balancing in golang:

type LoadBalancer struct {
	servers []string
	index   int
	mu      sync.Mutex
}

func (lb *LoadBalancer) getNextServer() string {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	server := lb.servers[lb.index]
	lb.index = (lb.index + 1) % len(lb.servers)
	return server
}

func (lb *LoadBalancer) handleRequest(w http.ResponseWriter, r *http.Request) {
	server := lb.getNextServer()
	fmt.Fprintf(w, "Redirecting to server: %s", server)
}

func main() {
	lb := &LoadBalancer{
		servers: []string{"http://server1", "http://server2", "http://server3"},
	}

	http.HandleFunc("/", lb.handleRequest)
	http.ListenAndServe(":8080", nil)
}

*****

Least Connections algorithm for load balancing in golang:

type Connection struct {
	server string
	count  int
}

type LoadBalancer struct {
	servers []Connection
	mu      sync.Mutex
}

func (lb *LoadBalancer) getNextServer() string {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	var leastIndex int
	for i, conn := range lb.servers {
		if conn.count < lb.servers[leastIndex].count {
			leastIndex = i
		}
	}
	lb.servers[leastIndex].count++
	return lb.servers[leastIndex].server
}

IP Hash algorithm for load balancing in golang:

func ipHash(ip string, servers []string) string {
	hash := 0
	for _, char := range ip {
		hash += int(char)
	}
	return servers[hash % len(servers)]
}


Random algorithm for load balancing in golang:

func getRandomServer(servers []string) string {
	return servers[rand.Intn(len(servers))]
}


Weighted Round Robin algorithm for load balancing in golang:

type WeightedServer struct {
	server string
	weight int
}

type LoadBalancer struct {
	servers []WeightedServer
	totalWeight int
}

func (lb *LoadBalancer) getNextServer() string {
	lb.mu.Lock()
	defer lb.mu.Unlock()
	r := rand.Intn(lb.totalWeight)
	for _, server := range lb.servers {
		if r < server.weight {
			return server.server
		}
		r -= server.weight
	}
	return ""
}


