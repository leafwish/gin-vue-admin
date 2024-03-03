import service from '@/utils/request'

export const createCollaboratorTeam = (data) => {
  return service({
    url: '/collaborator/team',
    method: 'post',
    data
  })
}

export const updateCollaboratorTeam = (data) => {
  return service({
    url: '/collaborator/team',
    method: 'put',
    data
  })
}

export const deleteCollaboratorTeam = (data) => {
  return service({
    url: '/collaborator/team',
    method: 'delete',
    data
  })
}

export const getCollaboratorTeam = (params) => {
  return service({
    url: '/collaborator/team',
    method: 'get',
    params
  })
}

export const getCollaboratorTeamPage = (params) => {
  return service({
    url: '/collaborator/team/page',
    method: 'get',
    params
  })
}
