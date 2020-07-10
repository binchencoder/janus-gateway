# Internal

这个目录下的文件原来在//gateway/runtime 下, 为了让runtime 目录下的文件更加清晰, 想将balancer.go、hook.go、service.go 迁移到//gateway/runtime 下, 但是在迁移的过程中遇到了难题

//gateway/rume  和 //internal 互相依赖造成死循环

# NOTE

这个目录目前没有用
