'use server'

import { token } from "@/actions/queries/auth/token"
import { env } from "@/lib/env"
import { Property } from "@/types/property"

type GetPropertyReponse = { properties: Property[] | undefined, status: number }

export async function listProperties(): Promise<GetPropertyReponse> {
  const auth_token = await token()

  try {
    const response = await fetch(`${env.SERVER_ADDR}/properties`, {
      method: "GET",
      headers: {
        "Authorization": `Bearer ${auth_token}`
      }
    })

    if (response.status !== 200) {
      return { properties: undefined, status: response.status }
    }

    const properties = await response.json() as Property[]

    return { properties, status: response.status }
  } catch {
    return { properties: undefined, status: 500 }
  }
}