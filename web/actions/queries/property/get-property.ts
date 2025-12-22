'use server'

import { token } from "@/actions/queries/auth/token"
import { env } from "@/lib/env"
import { Property } from "@/types/property"
import { Response } from "@/types/response"

type GetPropertyResponse = {
  property: Property | undefined,
  status: number
}

export async function getProperty(id: string): Promise<GetPropertyResponse> {
  const auth_token = await token()

  try {
    const response = await fetch(`${env.SERVER_ADDR}/property/${id}`, {
      method: "GET",
      headers: {
        "Authorization": `Bearer ${auth_token}`
      }
    })

    if (response.status !== 200) {
      return { property: undefined, status: response.status }
    }

    const { result: property } = await response.json() as Response<Property>

    return { property, status: response.status }
  } catch {
    return { property: undefined, status: 500 }
  }
}