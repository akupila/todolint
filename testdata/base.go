package base

// OK

// TODO(test): Make a pizza
// FIXME(test): Thinner crust
// BUG(test): Remove pineapple
/*
TODO: Multi-line comments are ignored
*/

// Error

// This is a TODO          // want `TODO should be at the beginning of the line`
// TODO(test)Make a pizza  // want `TODO should follow the format '// TODO\(context\): text'`
// TODO(test) Make a pizza // want `TODO should follow the format '// TODO\(context\): text'`
// TODO(test):Make a pizza // want `TODO should follow the format '// TODO\(context\): text'`
// TODO                    // want `TODO should include additional context: TODO\(<context>\)`
// TODO(): Make a pizza    // want `TODO should include additional context: TODO\(<context>\)`
// TODO: Make a pizza      // want `TODO should include additional context: TODO\(<context>\)`
// TODO(test):             // want `TODO should describe what needs to change`
// TODO(~): Make a pizza   // want `TODO context does not match regular expression: \\w+`
