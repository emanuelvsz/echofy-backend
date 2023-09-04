package cfg

type Strings struct {
	Dir string `env:"STRINGS_DIR" envDefault:"src/core/res/str/dict/pt-br"`
}
