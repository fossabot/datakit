#!/usr/bin/env python3
# -*- coding: utf-8 -*-

# This script copy from
#  https://github.com/DataDog/datadog-agent/blob/main/tasks/libs/copyright.py
# We just made some adjust according to specific conditions.

import re
import subprocess
import sys
import argparse
from pathlib import Path, PurePosixPath

GLOB_PATTERN = "**/*.go"

COPYRIGHT_HEADER = """
// Unless explicitly stated otherwise all files in this repository are licensed
// under the MIT License.
// This product includes software developed at TODO (https://www.datakit.tools/).
// Copyright TODO
""".strip()

COPYRIGHT_REGEX = [
    r'^// Unless explicitly stated otherwise all files in this repository are licensed$',
    r'^// under the MIT License\.$',
    r'^// This product includes software developed at TODO \(https://www\.datakit\.tools/\)\.$',
    r'^// Copyright TODO$',
]

# These path patterns are excluded from checks
PATH_EXCLUSION_REGEX = [
    # These are auto-generated files but without headers to indicate it
    '/vendor',
    '/plugins/externals/ebpf',
    '/plugins/inputs/skywalking/v3',
    '/internal/win_utils/pdh',
    '/internal/msgpack',
    '/.git/',
]

# These header matchers skip enforcement of the rules if found in the first
# line of the file
HEADER_EXCLUSION_REGEX = [
    '^// Code generated ',
    '^//go:generate ',
    '^// AUTOGENERATED FILE: ',
    '^// Copyright.* OpenTelemetry Authors',
    '^// Copyright.* The Go Authors',
    '^// This file includes software developed at CoreOS',
    '^// Copyright 2017 Kinvolk',
]


COMPILED_COPYRIGHT_REGEX = [re.compile(regex, re.UNICODE) for regex in COPYRIGHT_REGEX]
COMPILED_PATH_EXCLUSION_REGEX = [re.compile(regex, re.UNICODE) for regex in PATH_EXCLUSION_REGEX]
COMPILED_HEADER_EXCLUSION_REGEX = [re.compile(regex, re.UNICODE) for regex in HEADER_EXCLUSION_REGEX]


class CopyrightLinter:
    """
    This class is used to enforce copyright headers on specified file patterns
    """

    def __init__(self, debug=False):
        self._debug = debug

    @staticmethod
    def _get_repo_dir():
        script_dir = PurePosixPath(__file__).parent

        repo_dir = (
            subprocess.check_output(
                ['git', 'rev-parse', '--show-toplevel'],
                cwd=script_dir,
            )
            .decode(sys.stdout.encoding)
            .strip()
        )

        return PurePosixPath(repo_dir)

    @staticmethod
    def _is_excluded_path(filepath, exclude_matchers):
        for matcher in exclude_matchers:
            if re.search(matcher, filepath.as_posix()):
                return True

        return False

    @staticmethod
    def _get_matching_files(root_dir, glob_pattern, exclude=None):
        if exclude is None:
            exclude = []

        # Glob is a generator so we have to do the counting ourselves
        all_matching_files_cnt = 0

        filtered_files = []
        for filepath in Path(root_dir).glob(glob_pattern):
            all_matching_files_cnt += 1
            if not CopyrightLinter._is_excluded_path(filepath, exclude):
                filtered_files.append(filepath)

        excluded_files_cnt = all_matching_files_cnt - len(filtered_files)
        print(f"[INFO] Excluding {excluded_files_cnt} files based on path filters!")

        return sorted(filtered_files)

    @staticmethod
    def _get_header(filepath):
        header = []
        with open(filepath, "r") as file_obj:
            # We expect a specific header format which should be 4 lines
            for _ in range(4):
                header.append(file_obj.readline().strip())

        return header

    @staticmethod
    def _is_excluded_header(header, exclude=None):
        if exclude is None:
            exclude = []

        for matcher in exclude:
            if re.search(matcher, header[0]):
                return True

        return False

    def _has_copyright(self, filepath):
        header = CopyrightLinter._get_header(filepath)
        if header is None:
            print("[WARN] Mismatch found! Could not find any content in file!")
            return False

        if len(header) > 0 and CopyrightLinter._is_excluded_header(header, exclude=COMPILED_HEADER_EXCLUSION_REGEX):
            if self._debug:
                print(f"[INFO] Excluding {filepath} based on header '{header[0]}'")
            return True

        if len(header) <= 3:
            print("[WARN] Mismatch found! File too small for header stanza!")
            return False

        for line_idx, matcher in enumerate(COMPILED_COPYRIGHT_REGEX):
            if not re.match(matcher, header[line_idx]):
                print(
                    f"[WARN] Mismatch found! Expected '{COPYRIGHT_REGEX[line_idx]}' pattern but got '{header[line_idx]}'"
                )
                return False

        return True

    def _assert_copyrights(self, files):
        failing_files = []
        for filepath in files:
            if self._has_copyright(filepath):
                if self._debug:
                    print(f"[ OK ] {filepath}")

                continue

            print(f"[FAIL] {filepath}")
            failing_files.append(filepath)

        total_files = len(files)
        if failing_files:
            pct_failing = (len(failing_files) / total_files) * 100
            print()
            print(
                f"FAIL: There are {len(failing_files)} files out of "
                + f"{total_files} ({pct_failing:.2f}%) that are missing the proper copyright!"
            )

        return failing_files

    def _prepend_header(self, filepath, dry_run=True):
        with open(filepath, 'r+') as file_obj:
            existing_content = file_obj.read()

            if dry_run:
                return True

            file_obj.seek(0)
            new_content = COPYRIGHT_HEADER + "\n\n" + existing_content
            file_obj.write(new_content)

        # Verify result. A problem here is not benign so we stop the whole run.
        if not self._has_copyright(filepath):
            raise Exception(f"[ERROR] Header prepend failed to produce correct output for {filepath}!")

        return True

    @staticmethod
    def _is_build_header(line):
        return line.startswith("// +build ") or line.startswith("//+build ") or line.startswith("//go:build ")

    def _is_package_comment(line):
        return line.startswith("// Package ")

    def _fix_file_header(self, filepath, dry_run=True):
        header = CopyrightLinter._get_header(filepath)

        # Empty file - ignore
        if len(header) < 1:
            return False

        # If the file starts with a comment and it's not a build comment,
        # there is likely a manual fix to the header needed
        if header[0].startswith("//") and not CopyrightLinter._is_build_header(header[0]) and not CopyrightLinter._is_package_comment(header[0]):
            return False

        if dry_run:
            return True

        return self._prepend_header(filepath, dry_run=dry_run)

    def _fix(self, failing_files, dry_run=True):
        failing_files_cnt = len(failing_files)
        errors = []
        for idx, filepath in enumerate(failing_files):
            print(f"[INFO] ({idx+1:3d}/{failing_files_cnt:3}) Fixing '{filepath}'...")

            if not self._fix_file_header(filepath, dry_run=dry_run):
                error_message = f"'{filepath}' could not be fixed!"
                print(f"[WARN] ({idx+1:3d}/{failing_files_cnt:3}) {error_message}")
                errors.append(Exception(error_message))

        return errors

    def assert_compliance(self, fix=False, dry_run=True):
        """
        This method applies the GLOB_PATTERN to the root of the repository and
        verifies that all files have the expected copyright header.
        """
        git_repo_dir = CopyrightLinter._get_repo_dir()

        if self._debug:
            print(f"[DEBG] Repo root: {git_repo_dir}")
            print(f"[DEBG] Finding all files in {git_repo_dir} matching '{GLOB_PATTERN}'...")

        matching_files = CopyrightLinter._get_matching_files(
            git_repo_dir,
            GLOB_PATTERN,
            exclude=COMPILED_PATH_EXCLUSION_REGEX,
        )
        print(f"[INFO] Found {len(matching_files)} files matching '{GLOB_PATTERN}'")

        failing_files = self._assert_copyrights(matching_files)
        if len(failing_files) > 0:
            if not fix:
                print("CHECK: FAIL")
                raise Exception(
                    f"Copyright linting found {len(failing_files)} files that did not have the expected header!"
                )

            # If "fix=True", we will attempt to fix the failing files
            errors = self._fix(failing_files, dry_run=dry_run)
            if errors:
                raise Exception(f"Copyright linter was unable to fix {len(errors)}/{len(failing_files)} files!")

            return

        print("CHECK: OK")


if __name__ == '__main__':
    parser = argparse.ArgumentParser()
    parser.add_argument("--fix", dest="fix", action='store_true', help='auto add copyright to code')
    parser.add_argument("--dry-run", dest="dry_run", action='store_true', help='dry run')

    args = parser.parse_args()
    #CopyrightLinter(debug=True).assert_compliance(fix=True, dry_run=False)

    print(args)

    CopyrightLinter(debug=True).assert_compliance(fix=args.fix, dry_run=args.dry_run)
