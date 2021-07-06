package global

// Init Global Init Entrance
func Init() error {
	var err error
	// TODO 初始化配置文件
	if err = InitKubernetesClient(); err != nil {
		return err
	}
	return nil
}