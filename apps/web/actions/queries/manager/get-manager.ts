'use server'

import { env } from "@/lib/env"
import { Manager } from "@/types/manager"
import { Response } from "@/types/response"
import { token } from "../auth/token"

type GetManagerResponse = {
  manager: Manager | undefined,
  status: number
}

export async function getManager(): Promise<GetManagerResponse> {
  const auth_token = await token()

  if (!auth_token) {
    return { manager: undefined, status: 500 }
  }

  try {
    const response = await fetch(`${env.SERVER_ADDR}/manager`, {
      method: "GET",
      headers: {
        'Authorization': `Bearer ${auth_token}`
      }
    })

    if (response.status !== 200) {
      return { manager: undefined, status: response.status }
    }

    const { result: manager } = await response.json() as Response<Manager>

    return { manager, status: response.status }
  } catch {
    return { manager: undefined, status: 500 }
  }
}