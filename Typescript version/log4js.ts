import * as log4js from 'log4js'
import { INFO } from './settings'
const log = console.log
type _Logs = {
	(temp: string): void
}

log4js.configure({
	appenders: {
		Notice: { type: 'file', filename: 'logs/notice.log', compress: true },
	},
	categories: {
		default: { appenders: ['Notice'], level: 'debug' },
	},
})
export const Logger = log4js.getLogger('Notice')

export const log4Info: _Logs = (info) => {
	INFO.Logger && log(info)
}
