import { Address } from "./address";

export type Tenant = {
  id: string
  manager_id: string
  fullname: string
  cpf: string
  rg: string
  phone: string
  occupation: string
  marital_status: 'Solteiro(a)' | 'Casado(a)' | 'Amasiado(a)' | 'Divorciado(a)' | 'Viúvo(a)' | 'União Estável',
  address: Address,
};


