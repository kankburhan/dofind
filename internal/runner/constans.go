package runner

const (
	version = "1.2.0"
	header  = `
	██████╗░░█████╗░███████╗██╗███╗░░██╗██████╗░
	██╔══██╗██╔══██╗██╔════╝██║████╗░██║██╔══██╗
	██║░░██║██║░░██║█████╗░░██║██╔██╗██║██║░░██║
	██║░░██║██║░░██║██╔══╝░░██║██║╚████║██║░░██║
	██████╔╝╚█████╔╝██║░░░░░██║██║░╚███║██████╔╝
	╚═════╝░░╚════╝░╚═╝░░░░░╚═╝╚═╝░░╚══╝╚═════╝░
	______________ by @kankburhan ______________
	`
	usage = `
	dofind [options]`
	options = `
	-t, --target <TARGET>		Target find domain
	-o, --output <FILE>		File to save results
	-c, --concurrent <i>		Set the concurrency level (default: 20)
	-s, --silent			Silent mode
	-v, --verbose			Verbose mode
	-u, --update-domain		Update domain list
	-h, --help			Display its help
	`
	Timeout    = 1000
	pathfile   = "/.config/.dofind/"
	configfile = "config.json"
)
