export interface ClockData {
  task: string
  parents: string
  category: string
  start: string
  end: string
  effort: string
  ishabit: string
  tags: string[]
}


export function isClockData(arg: any): arg is ClockData {
  return arg.task !== undefined
    && arg.parents !== undefined
    && arg.category !== undefined
    && arg.start !== undefined
    && arg.end !== undefined
    && arg.effort !== undefined
    && arg.ishabit !== undefined
    && arg.tags !== undefined
    && Array.isArray(arg.tags)
}
