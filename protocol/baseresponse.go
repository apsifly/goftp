package protocol

//standard client breaks on multiline messages
type baseResponse struct {
	code    string
	message string
}

var Response110 = baseResponse{
	code: "110",
	message: `Restart marker reply.
`,
}
var Response120 = baseResponse{
	code: "120",
	message: `Service ready in nnn minutes.
`,
}
var Response125 = baseResponse{
	code: "125",
	message: `Data connection already open; transfer starting.
`,
}
var Response150 = baseResponse{
	code: "150",
	message: `File status okay; about to open data connection.
`,
}
var Response200 = baseResponse{
	code: "200",
	message: `Command okay.
`,
}
var Response202 = baseResponse{
	code: "202",
	message: `Command not implemented, superfluous at this site.
`,
}
var Response211 = baseResponse{
	code: "211",
	message: `System status, or system help reply.
`,
}
var Response212 = baseResponse{
	code: "212",
	message: `Directory status.
`,
}
var Response213 = baseResponse{
	code: "213",
	message: `File status.
`,
}
var Response214 = baseResponse{
	code: "214",
	message: `Help message.
`,
}
var Response215 = baseResponse{
	code: "215",
	message: `NAME system type.
`,
}
var Response220 = baseResponse{
	code: "220",
	message: `Service ready for new user.
`,
}
var Response221 = baseResponse{
	code: "221",
	message: `Service closing control connection. Logged out if appropriate.
`,
}
var Response225 = baseResponse{
	code: "225",
	message: `Data connection open; no transfer in progress.
`,
}
var Response226 = baseResponse{
	code: "226",
	message: `Closing data connection.
`,
}
var Response227 = baseResponse{
	code: "227",
	message: `Entering Passive Mode (h1,h2,h3,h4,p1,p2).
`,
}
var Response230 = baseResponse{
	code: "230",
	message: `User logged in, proceed.
`,
}
var Response250 = baseResponse{
	code: "250",
	message: `Requested file action okay, completed.
`,
}
var Response257 = baseResponse{
	code: "257",
	message: `"PATHNAME" created.

`,
}
var Response331 = baseResponse{
	code: "331",
	message: `User name okay, need password.
`,
}
var Response332 = baseResponse{
	code: "332",
	message: `Need account for login.
`,
}
var Response350 = baseResponse{
	code: "350",
	message: `Requested file action pending further information.

`,
}
var Response421 = baseResponse{
	code: "421",
	message: `Service not available, closing control connection.
`,
}
var Response425 = baseResponse{
	code: "425",
	message: `Can't open data connection.
`,
}
var Response426 = baseResponse{
	code: "426",
	message: `Connection closed; transfer aborted.
`,
}
var Response450 = baseResponse{
	code: "450",
	message: `Requested file action not taken. File unavailable.
`,
}
var Response451 = baseResponse{
	code: "451",
	message: `Requested action aborted: local error in processing.
`,
}
var Response452 = baseResponse{
	code: "452",
	message: `Requested action not taken. Insufficient storage space in system.
`,
}
var Response500 = baseResponse{
	code: "500",
	message: `Syntax error, command unrecognized.
`,
}
var Response501 = baseResponse{
	code: "501",
	message: `Syntax error in parameters or arguments.
`,
}
var Response502 = baseResponse{
	code: "502",
	message: `Command not implemented.
`,
}
var Response503 = baseResponse{
	code: "503",
	message: `Bad sequence of commands.
`,
}
var Response504 = baseResponse{
	code: "504",
	message: `Command not implemented for that parameter.
`,
}
var Response530 = baseResponse{
	code: "530",
	message: `Not logged in.
`,
}
var Response532 = baseResponse{
	code: "532",
	message: `Need account for storing files.
`,
}
var Response550 = baseResponse{
	code: "550",
	message: `Requested action not taken. File unavailable.
`,
}
var Response551 = baseResponse{
	code: "551",
	message: `Requested action aborted: page type unknown.
`,
}
var Response552 = baseResponse{
	code: "552",
	message: `Requested file action aborted. Exceeded storage allocation.
`,
}
var Response553 = baseResponse{
	code: "553",
	message: `Requested action not taken. File name not allowed.
`,
}
