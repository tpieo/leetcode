func findRedundantConnection(edges [][]int) []int {
    // n 个节点相连, n = len(edges) 提示中说明
    n := len(edges)
    // p[i] 第 i 个节点的祖先 
    p := make([]int, n+1)
    // 每个节点的祖先初始化为自己
    for i := 1; i < n+1; i++ {
        p[i] = i
    }
    // 查找 x 节点祖先的函数
    find := func(x int) int {
        for p[x] != x {
            p[x] = p[p[x]]  //路径压缩，想象一下图
            x = p[x]
        }
        return x
    }
    // 连接 x, y 节点的函数
    connect := func(x, y int) {
        p[find(x)] = find(y)
    }
    // 判断 x, y 节点是否相连
    isConnected := func(x, y int) bool {
        return find(x) == find(y)
    }
    for i := 0; i < n; i++ {
        x, y := edges[i][0], edges[i][1]
        if isConnected(x, y) {
            return edges[i]
        }
        connect(x, y)
    }
    return edges[0]
}