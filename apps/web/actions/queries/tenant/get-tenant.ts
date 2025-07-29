'use server'

import { token } from "@/actions/queries/auth/token"
import { env } from "@/lib/env"
import { Response } from "@/types/response"
import { Tenant } from "@/types/tenant"

type GetTenantResponse = {
  tenant: Tenant | undefined,
  status: number
}

export async function getTenant(id: string): Promise<GetTenantResponse> {
  const auth_token = await token()

  try {
    const response = await fetch(`${env.SERVER_ADDR}/tenant/${id}`, {
      method: "GET",
      headers: {
        "Authorization": `Bearer ${auth_token}`
      }
    })

    if (response.status != 200) {
      return { tenant: undefined, status: response.status }
    }

    const { result: tenant } = await response.json() as Response<Tenant>

    return { tenant, status: response.status }
  } catch {
    return { tenant: undefined, status: 500 }
  }
}