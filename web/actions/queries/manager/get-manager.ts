"use server"

import { env } from "@/lib/env"
import type { Manager } from "@/types/manager"
import type { Response } from "@/types/response"
import { token } from "../auth/token"

interface GetManagerResponse {
	manager: Manager | undefined
}

export async function getManager(): Promise<GetManagerResponse> {
	const authToken = await token()

	if (!authToken) {
		return { manager: undefined }
	}

	try {
		const response = await fetch(`${env.SERVER_ADDR}/manager`, {
			method: "GET",
			headers: {
				Authorization: `Bearer ${authToken}`,
			},
		})

		if (response.status !== 200) {
			return { manager: undefined }
		}

		const { result: manager } = (await response.json()) as Response<Manager>

		return { manager }
	} catch {
		return { manager: undefined }
	}
}
