export type CEP = {
  cep: string
  state: string
}

export interface CepService {
  SearchCEP(cep: string): Promise<CEP | undefined>
}