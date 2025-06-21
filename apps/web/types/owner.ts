export type Owner = {
  id: string
  manager_id: string
  fullname: string
  cpf: string
  rg: string
  email: string
  phone: string
  occupation: string
  marital_status: 'Solteiro(a)' | 'Casado(a)' | 'Amasiado(a)' | 'Divorciado(a)' | 'Viúvo(a)' | 'União Estável',
  address: {
    address: string
    street: string
    number: string
    neighborhood: string
    complement: string
    city: string
    state: string
    state_abbr: string
    zip_code: string
  },
};


