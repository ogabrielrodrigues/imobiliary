export type CEP = {
  street: string
  cep: string
  state: string
  city: string
  neighborhood: string
}

export interface CepService {
  SearchCEP(cep: string): Promise<CEP | undefined>
}