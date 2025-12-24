"use server"

import { decodeJWT } from "@/lib/jwt"
import { token } from "./token"

export async function auth(): Promise<string | undefined> {
	const authToken = await token()

	if (!authToken) {
		return undefined
	}

	return await decodeJWT(authToken)
}
