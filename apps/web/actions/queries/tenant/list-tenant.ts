'use server'

import { token } from "@/actions/queries/auth/token"
import { env } from "@/lib/env"
import { Response } from "@/types/response"
import { Tenant } from "@/types/tenant"

type ListTenantsResponse = {
  tenants: Tenant[] | undefined,
  status: number
}

export async function listTenants(): Promise<ListTenantsResponse> {
  const auth_token = await token()

  try {
    const response = await fetch(`${env.SERVER_ADDR}/tenant`, {
      method: "GET",
      headers: {
        "Authorization": `Bearer ${auth_token}`
      }
    })

    if (response.status != 200) {
      return { tenants: undefined, status: response.status }
    }

    const { result: tenants } = await response.json() as Response<Tenant[]>

    return { tenants, status: response.status }
  } catch {
    return { tenants: undefined, status: 500 }
  }
}