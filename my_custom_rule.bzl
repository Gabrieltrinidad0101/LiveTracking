def _write_new_file_impl(ctx):
    input_file = ctx.file.my_input_file
    output_file = ctx.actions.declare_file(ctx.attr.my_output_file + ".txt")

    command = 'cat "{}" > "{}"'.format(input_file.path, output_file.path)

    ctx.actions.run(
        inputs = [input_file],
        outputs = [output_file],
        executable = "/bin/bash",
        arguments = ["-c", command],
    )

    return [DefaultInfo(files = depset([output_file]))]

write_new_file = rule(
    implementation = _write_new_file_impl,
    attrs = {
        "my_input_file": attr.label(allow_single_file = True),
        "my_output_file": attr.string(),
    }
)
