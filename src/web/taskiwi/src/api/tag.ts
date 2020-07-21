import client from "./client";
import { ALL_TAGS_PATH } from "../constants/api_endpoints";
import { AxiosPromise } from "axios";
import { Tags } from "../model/Tags";

export const fetchTags = (): AxiosPromise<Tags> => client.get(ALL_TAGS_PATH)
