'use server'

import { env } from "@/lib/env"
import { token } from "./auth"
import { Plan } from "@/types/plan"

export async function getPlan() {
  const auth_token = await token()

  const response = await fetch(`${env.SERVER_ADDR}/users/plan`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${auth_token}`
    }
  })

  const plan = await response.json() as Plan

  return plan
}