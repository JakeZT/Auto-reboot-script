import dotenv from 'dotenv'
import fs from 'fs'
import path from 'path'

const config = dotenv.parse(fs.readFileSync(path.join(__dirname, '.env')))
for (const k in config) {
	process.env[k] = config[k]
}

export const Tasks = ['p', 'p3']
export const Test = true
export const Rules = {
	rule1: '10 22 1 * * *',
	rule2: { hour: 22, minute: 10 },
}
export const EMAIL: EmailType = {
	Email_Open: Boolean(process.env.EMAIL_OPEN),
	Email_Address: process.env.FROM || '',
	Email_Service: process.env.SERVICE || '',
	Smtp_Pass: process.env.SMTP_PASS || '',
	Email_To: process.env.TO || '',
}
type EmailType = {
	Email_Open: boolean
	Email_Address: string
	Email_Service: string
	Smtp_Pass: string
	Email_To: string
}

export const INFO = {
	Logger: Boolean(process.env.LOGGER),
	Name: 'FleetingS',
}
