import * as nodemailer from 'nodemailer'
import { EMAIL } from './settings'
import { Logger } from './log4js'
type _Mail = (username: string) => void

export const Mail: _Mail = async (username = 'XXXX') => {
	let transporter = nodemailer.createTransport({
		port: 465,
		service: EMAIL.Email_Service,
		secure: true,
		auth: {
			user: EMAIL.Email_Address,
			pass: EMAIL.Smtp_Pass,
		},
	})

	let info = await transporter.sendMail({
		from: `"Reboot engine works well" <${EMAIL.Email_Address}>`,
		to: EMAIL.Email_To,
		subject: `${username} stays online today: ${new Date().toLocaleDateString()}`, // Subject line
		text: ``,
	})
	Logger.info(`Message sent at ${new Date().toISOString()}, ID is ${info.messageId}`)
	console.log('Message sent: %s', info.messageId)
}
export default Mail
