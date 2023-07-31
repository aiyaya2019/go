package ziface

/*
连接管理模块抽象层
*/
type IConnManager interface {
	//添加连接
	Add(conn IConnectioin)

	//删除连接
	Remove(conn IConnectioin)

	//根据ConnID获取连接
	Get(connId uint32) (IConnectioin, error)

	//得到当前连接总数
	Len() int

	//清除并终止所有连接
	ClearConn()
}
