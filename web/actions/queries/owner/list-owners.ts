'use server'

import { token } from "@/actions/queries/auth/token"
import { env } from "@/lib/env"
import { Owner } from "@/types/owner"
import { Response } from "@/types/response"

type ListOwnersResponse = {
  owners: Owner[] | undefined,
  status: number
}

export async function listOwners(): Promise<ListOwnersResponse> {
  const auth_token = await token()

  try {
    const response = await fetch(`${env.SERVER_ADDR}/owners`, {
      method: "GET",
      headers: {
        "Authorization": `Bearer ${auth_token}`
      }
    })

    if (response.status != 200) {
      return { owners: undefined, status: response.status }
    }

    const { result: owners } = await response.json() as Response<Owner[]>

    return { owners, status: response.status }
  } catch {
    return { owners: undefined, status: 500 }
  }
}