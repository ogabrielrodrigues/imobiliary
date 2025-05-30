'use server'

import { token } from "@/actions/queries/auth/token"
import { OwnerRequest } from "@/app/dashboard/locacao/proprietarios/_components/new-owner-form"
import { searchCEP } from "@/lib/cep/api-brasil"
import { env } from "@/lib/env"

export async function createOwner(data: OwnerRequest): Promise<number> {
  const auth_token = await token()

  const found = await searchCEP(data.address.zip_code)

  if (found == undefined) {
    return 400
  }

  data.address.state = found.state

  try {
    const response = await fetch(`${env.SERVER_ADDR}/owners`, {
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