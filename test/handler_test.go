package test

import (
	"magicTGArchive/internal/pkg/env"
	"testing"
)

//FIXME -> only works if viper.SetConfigFile("../.env") why?

func TestReceiveEnvVars(t *testing.T) {
	conf, err := env.ReceiveEnvVars()
	if err != nil {
		t.Fatalf("ReceiveEnvVars throw error: %v", err)
	}

	if conf.DbUser == ""{
		t.Error("conf.DbUser shouldn't be empty")
	}else if conf.DbPass == ""{
		t.Error("conf.DbPass shouldn't be empty")
	}else if conf.DbPort == ""{
		t.Error("conf.DbPort shouldn't be empty")
	}else if conf.DbName == ""{
		t.Error("conf.DbName shouldn't be empty")
	}else if conf.DbCollAllCards == ""{
		t.Error("conf.DbCollAllCards shouldn't be empty")
	}else if conf.DbCollMyCards == ""{
		t.Error("conf.DbCollMyCards shouldn't be empty")
	}else if conf.DbCollImgInfo == ""{
		t.Error("conf.DbCollImgInfo shouldn't be empty")
	}

}
