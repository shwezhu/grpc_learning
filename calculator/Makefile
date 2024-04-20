# run: make generate
generate:
	protoc --proto_path=proto proto/*.proto --go_out=. --go-grpc_out=.

# 1. 上面使用的所有路径都是相对 protoc 命令执行时所在的目录

# 2. --proto_path=proto: 指定 .proto 文件的搜索路径, 在这，编译器会在名为 proto 的目录下查找 .proto 文件,
# 你也可以使用 -IPATH 命令行参数来指定搜索路径 如 protoc -I. --go_out=. --go-grpc_out=. proto/*.proto
# proto/*.proto : 指定要编译的 .proto 文件，这里我们指定了 proto 目录下的所有 .proto 文件,
# 不指定会报错, 因为 protoc 需要知道要编译哪个文件, 你可以使用如 proto/xxx.proto

# 3. --go_out=.: 还记得我们在 .proto 文件中指定了 go_package="./pb" 吗,
# --go_out=. 为 go_package 指定了所在的工作目录, 举例:
# 若你使用 --go_out=abc, go_package="./pb", 则生成的文件会放在 abc/pb 目录下, 前提是你得先创建 abc 目录
# 我们在这指定的是 --go_out=., 所以最终生成的文件会放到 ./pb 目录下