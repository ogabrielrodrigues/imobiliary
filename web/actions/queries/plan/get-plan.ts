'use server'

import { env } from "@/lib/env"
import { Plan } from "@/types/plan"
import { token } from "../token"

export async function getPlan(): Promise<{ status: number, plan: Plan | undefined }> {
  const auth_token = await token()

  try {
    const response = await fetch(`${env.SERVER_ADDR}/users/plan`, {
      method: "GET",
      headers: {
        "Authorization": `Bearer ${auth_token}`
      }
    })

    return { status: response.status, plan: await response.json() }
  } catch {
    return { status: 500, plan: undefined }
  }
}