package util

import (
	"os"
	"reflect"
	"strconv"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func BindFromJSON(dest any, filename, path string) error {
	v := viper.New()
	v.SetConfigType("json")
	v.AddConfigPath(path)
	v.SetConfigName(filename)

	err := v.ReadInConfig()

	if err != nil {
		return err
	} else {
		err = v.Unmarshal(&dest)

		if err != nil {
			logrus.Error("Failed to unmarshal config: ", err)
			return err
		}
	}
	return nil
}

func SetEnvFromConsulKV(v *viper.Viper) error {
	env := make(map[string]string)

	err := v.Unmarshal(&env)

	if err != nil {
		logrus.Error("Failed to unmarshal config: ", err)
		return err
	}

	for k, v := range env {
		var (
			valOf = reflect.ValueOf(v)
			val   string
		)
		switch valOf.Kind() {
		case reflect.String:
			val = valOf.String()
		case reflect.Int:
			val = strconv.Itoa(int(valOf.Int()))
		case reflect.Uint:
			val = strconv.Itoa(int(valOf.Uint()))
		case reflect.Float32:
			val = strconv.Itoa(int(valOf.Float()))
		case reflect.Float64:
			val = strconv.Itoa(int(valOf.Float()))
		case reflect.Bool:
			val = strconv.FormatBool(valOf.Bool())
		default:
			panic("unsupported type")
		}
		err = os.Setenv(k, val)

		if err != nil {
			logrus.Error("Failed to set env: ", err)
			return err
		}
	}
	return nil
}

func BindFromConsul(dest any, endpoint, path string) error {
	v := viper.New()
	v.SetConfigType("json")
	err := v.AddRemoteProvider("consul", endpoint, path)

	if err != nil {
		logrus.Error("Failed to add remote provider: ", err)
		return err
	}

	err = v.ReadRemoteConfig()
	if err != nil {
		logrus.Error("Failed to read remote config: ", err)
		return err
	}

	err = v.Unmarshal(&dest)
	if err != nil {
		logrus.Error("Failed to unmarshal/decode config: ", err)
		return err
	}

	err = SetEnvFromConsulKV(v)
	if err != nil {
		logrus.Error("Failed to set env from consul kv: ", err)
		return err
	}
	return nil
}
