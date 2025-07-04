export type Property = {
  id: string
  water_id: string
  energy_id: string
  user_id: string
  owner_id: string
  kind: 'Residencial' | 'Comercial' | 'Industrial' | 'Terreno' | 'Rural',
  status: 'Disponível' | 'Ocupado' | 'Indisponível' | 'Reservado' | 'Reformando',
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

