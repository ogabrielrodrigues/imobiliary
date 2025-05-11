'use server'

import { env } from "../env"
import { CEP } from "./cep"

export async function searchCEP(cep: string): Promise<CEP | undefined> {
  const response = await fetch(`${env.CEP_API_ADDR}/${cep}`)

  if (response.status != 200) {
    return undefined
  }

  const data = await response.json()

  const found_cep: CEP = {
    cep: data.cep,
    state: data.state
  }

  return found_cep
}