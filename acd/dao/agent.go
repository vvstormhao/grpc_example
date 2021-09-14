package dao

func GetAgentState(appId, agentId string) (uint32, error) {
	/*
		var agent *model.Agent
		result := mysqlconns.Instance().DB().Model(&model.Agent{}).Select("state").Where("app_id = ?", appId).Where("agent_id = ?", agentId).First(agent)
		if nil == result.Error {
			return uint32(agent.State), nil
		}

		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return 0, ErrRecordNotFound
		}

		return 0, ErrDBRequestError
	*/

	return 2, nil
}
