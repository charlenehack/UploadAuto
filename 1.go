1. 删除单个元素
func delItem(vs []string, s string) []string{
    for i := 0; i < len(vs); i++ {
          if s == vs[i] {
                vs = append(vs[:i], vs[i+1:]...)
                  i = i-1
          }
    }
    return vs
}
2.删除多个元素
func delItems(vs []string, dels []string) []string {
      dMap := make(map[string]bool)
      for _, s := range dels {
            dMap[s] = true
      }

      for i := 0; i < len(vs); i++ {
            if _, ok := dMap[vs[i]]; ok {
                  vs = append(vs[:i], vs[i+1:]...)
                  i = i-1
            }
      }
      return vs
}
