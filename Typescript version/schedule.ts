import * as shell from 'shelljs'
import schedule from 'node-schedule'
import { Tasks, Rules, INFO, Test } from './settings'
import { Logger } from './log4js'
import Mail from './mailer'
const sleep = async (time = 2000) => new Promise((resolve) => setTimeout(() => resolve(true), time))

const createCommandQueue = (task: string): string[] => {
	const basicPath = `/home/${task}`
	return [`cd ${basicPath}`, `pm2 stop ${task}`, `rimraf ${basicPath}/logs`, `pm2 start ${basicPath}/auto.js -n ${task}`]
}
async function rebootEngine(task: string) {
	try {
		const cmdQueue = createCommandQueue(task)
		for (let command of cmdQueue) {
			await runCommand(command)
		}
		Logger.info(`Successfully started ${task} --${new Date().toISOString()}`)
		await sleep(5000)
	} catch (err: unknown) {
		Logger.info(`Error: execute command failed --${new Date().toISOString()}`)
		shell.echo('Error: execute command failed ')
		shell.exit(1)
	}
}

const runTask = async () => {
	for (let task of Tasks) {
		await rebootEngine(task)
	}
}

const runCommand = (cmd: string) =>
	new Promise(async (resolve, reject) => {
		const res = shell.exec(cmd)
		if (res.code !== 0) {
			reject('Failed to execute command.')
		} else {
			await sleep()
			resolve(true)
		}
	})

const main = () => {
	return schedule.scheduleJob(Rules.rule1, () => {
		console.log('start...')
		runTask()
		Mail(INFO.Name)
		Logger.info(`successfully executed today. --${new Date().toISOString()}`)
		console.log('successfully executed today. ')
	})
}
main()
