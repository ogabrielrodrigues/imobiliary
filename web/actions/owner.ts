'use server'

import { Owner } from "@/types/owner"
import { token } from "./auth"
import { env } from "@/lib/env"
import { z } from "zod"
import { owner_schema } from "@/app/dashboard/locacao/proprietarios/_components/new-owner-form"
import { searchCEP } from "@/lib/cep/api-brasil"
import { assign_property_schema } from "@/app/dashboard/locacao/imoveis/[property_id]/_components/assign-property-form"

export async function getOwner(owner_id: string): Promise<{ status: number, owner: Owner }> {
  const auth_token = await token()

  if (!owner_id) {
    return { status: 400, owner: {} as Owner }
  }

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

  const owners = await response.json() as Owner[]


  return { status: response.status, owners: owners || [] }
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

export async function assignOwner(property_id: string, data: z.infer<typeof assign_property_schema>): Promise<number> {
  const auth_token = await token()

  const body = {
    ...data,
    property_id
  }

  const response = await fetch(`${env.SERVER_ADDR}/owners/assign`, {
    method: "PUT",
    body: JSON.stringify(body),
    headers: {
      "Authorization": `Bearer ${auth_token}`,
    }
  })

  return response.status
}