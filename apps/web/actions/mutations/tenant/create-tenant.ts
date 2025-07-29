'use server'

import { token } from "@/actions/queries/auth/token"
import { TenantRequest } from "@/app/dashboard/locacao/inquilinos/_components/new-tenant-form"
import { searchCEP } from "@/lib/cep/api-brasil"
import { env } from "@/lib/env"

export async function createTenant(data: TenantRequest): Promise<number> {
  const auth_token = await token()

  const found = await searchCEP(data.address.zip_code)

  if (found == undefined) {
    return 400
  }

  data.address.state = found.state

  try {
    const response = await fetch(`${env.SERVER_ADDR}/tenant`, {
      method: "POST",
      body: JSON.stringify(data),
      headers: {
        "Authorization": `Bearer ${auth_token}`,
      },
    })

    return response.status
  } catch {
    return 500
  }
}