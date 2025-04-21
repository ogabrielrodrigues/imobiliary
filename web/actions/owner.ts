'use server'

import { Owner } from "@/types/owner"
import { token } from "./auth"
import { env } from "@/lib/env"
import { z } from "zod"
import { owner_schema } from "@/app/dashboard/locacao/proprietarios/_components/new-owner-form"
import { searchCEP } from "@/lib/cep/api-brasil"

export async function getOwner(owner_id: string): Promise<{ status: number, owner: Owner }> {
  const auth_token = await token()

  const response = await fetch(`${env.SERVER_ADDR}/owners/${owner_id}`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${auth_token}`
    }
  })

  const owner = await response.json() as Owner

  return { status: response.status, owner }
}

export async function getOwners(): Promise<{ status: number, owners: Owner[] }> {
  const auth_token = await token()

  const response = await fetch(`${env.SERVER_ADDR}/owners`, {
    method: "GET",
    headers: {
      "Authorization": `Bearer ${auth_token}`
    }
  })

  const properties = await response.json() as Owner[]

  return { status: response.status, owners: properties }
}

export async function createOwner(data: z.infer<typeof owner_schema>): Promise<number> {
  const auth_token = await token()

  const cep = data.address.zip_code.replace("-", "")
  const found = await searchCEP(env.CEP_API_ADDR!, cep)

  if (found == undefined) {
    return 400
  }

  data.address.state = found.state

  const response = await fetch(`${env.SERVER_ADDR}/owners`, {
    method: "POST",
    body: JSON.stringify(data),
    headers: {
      "Authorization": `Bearer ${auth_token}`,
    },
  })

  return response.status
}