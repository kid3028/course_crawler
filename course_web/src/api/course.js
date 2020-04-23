import {Service} from "../lib/http";

export function dayStatistic(recordTime, current, size) {
    return Service({
        url: "/v1/dayStatistic",
        params : {"recordTime": recordTime, "current": current, "size": size}
    })
}

export function courseList(recordTime, subject, current, size) {
    return Service({
        url: "/v1/subjectCourseList",
        params : {"recordTime": recordTime, "subject": subject, "current": current, "size": size}
    })
}

export function courseDetail(cid) {
    return Service({
        url: "/v1/courseDetail",
        params : {cid: cid}
    })

}