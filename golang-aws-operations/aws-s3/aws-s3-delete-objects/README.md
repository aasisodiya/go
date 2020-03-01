# Delete Objects in S3 Bucket

## Question

* Why doesn't it return error if we try to delete object that doesn't exists?

    **Answer:** Because that's what the specs says it should do -> "Removes the null version (if there is one) of an object and inserts a delete marker, which becomes the latest version of the object. If there isn't a null version, Amazon S3 does not remove any objects.", if the object doesn't exists, it's still not an error when calling deleteObject

* `Error: cannot use objectkeys[i] (type string) as type *string in field value`

    **Answer:** The AWS SDK uses string pointers for most of its inputs due to the vagueries of the AWS API. It does include a helper function, aws.String, for this purpose: aws.String("STRING"),

* `Error: invalid indirect of s3.ObjectIdentifier literal (type s3.ObjectIdentifier)`

  **Answer:** You'll need to create a pointer to the value you're creating, which is done with & , * does the opposite, it dereferences a pointer.

  ```golang
  var objects []*s3.ObjectIdentifier
  for i := 0; i < len(objectkeys); i++ {
      objects[i] = &s3.ObjectIdentifier{
          Key: aws.String(objectkeys[i]),
      }
  }
  ```

## Reference

* [Why does S3.deleteObject not fail when the specified key doesn't exist?
](https://stackoverflow.com/questions/30697746/why-does-s3-deleteobject-not-fail-when-the-specified-key-doesnt-exist)
* [Why is *a{…} invalid indirect?
](https://stackoverflow.com/questions/20890850/why-is-a-invalid-indirect)
