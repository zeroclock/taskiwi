import client from './client'
import { ALL_TAGS_PATH } from '../constants/api_endpoints'
import { AxiosPromise } from 'axios'
import { Tags } from '../model/Tags'
import { TagsReq } from '../interface/request'

export const fetchTags = (params: TagsReq): AxiosPromise<Tags> =>
  client.post(ALL_TAGS_PATH, params)
