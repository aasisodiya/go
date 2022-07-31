# Delete Object in S3 Bucket

![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go.golang-aws-operations.aws-s3.delete-object&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)

## Question

- Why doesn't it return error if we try to delete object that doesn't exists?

  **Answer:** Because that's what the specs says it should do -> "Removes the null version (if there is one) of an object and inserts a delete marker, which becomes the latest version of the object. If there isn't a null version, Amazon S3 does not remove any objects.", if the object doesn't exists, it's still not an error when calling deleteObject

## Reference

- [Why does S3.deleteObject not fail when the specified key doesn't exist?
  ](https://stackoverflow.com/questions/30697746/why-does-s3-deleteobject-not-fail-when-the-specified-key-doesnt-exist)

[![Visitors](https://api.visitorbadge.io/api/visitors?path=aasisodiya.go&label=aasisodiya/go&labelColor=%23ffa500&countColor=%23263759&labelStyle=upper)](https://visitorbadge.io/status?path=aasisodiya.go)
