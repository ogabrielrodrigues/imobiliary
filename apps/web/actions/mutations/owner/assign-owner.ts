'use server'

import { token } from "@/actions/queries/auth/token"
import { AssignOwnerRequestData } from "@/app/dashboard/locacao/imoveis/[property_id]/_components/assign-property-form"
import { env } from "@/lib/env"

type AssignOwnerRequest = {
  id: string,
  data: AssignOwnerRequestData
}

export async function assignOwner({ id, data }: AssignOwnerRequest): Promise<number> {
  const auth_token = await token()

  const body = {
    property_id: id,
    owner_id: data.owner_id
  }

  try {
    const response = await fetch(`${env.SERVER_ADDR}/owners/assign`, {
      method: "PUT",
      body: JSON.stringify(body),
      headers: {
        "Authorization": `Bearer ${auth_token}`,
      }
    })

    return response.status
  } catch {
    return 500
  }
}