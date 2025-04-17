'use server'

import { CEP } from "./cep"

export async function searchCEP(api_addr: string, cep: string): Promise<CEP | undefined> {
  const response = await fetch(`${api_addr}/${cep}`)

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