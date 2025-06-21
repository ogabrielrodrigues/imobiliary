'use server'

import { token } from "@/actions/queries/auth/token"
import { env } from "@/lib/env"
import { Owner } from "@/types/owner"
import { Response } from "@/types/response"

type GetOwnerResponse = {
  owner: Owner | undefined,
  status: number
}

export async function getOwner(id: string): Promise<GetOwnerResponse> {
  const auth_token = await token()

  try {
    const response = await fetch(`${env.SERVER_ADDR}/owner/${id}`, {
      method: "GET",
      headers: {
        "Authorization": `Bearer ${auth_token}`
      }
    })

    if (response.status != 200) {
      return { owner: undefined, status: response.status }
    }

    const { result: owner } = await response.json() as Response<Owner>

    return { owner, status: response.status }
  } catch {
    return { owner: undefined, status: 500 }
  }
}