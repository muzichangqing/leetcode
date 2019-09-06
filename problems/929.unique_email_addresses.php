<?php
class Solution {

    /**
     * @param String[] $emails
     * @return Integer
     */
    function numUniqueEmails($emails) {
        $uniqueEmails = [];
        foreach($emails as $email) {
            $list = explode('@', $email);
            $domain = $list[1];
            $localName = $list[0];
            $localName = explode('+', $localName)[0];
            $localName = str_replace('.', '', $localName);
            $uniqueEmail = $localName . '@' . $domain;
            if (!in_array($uniqueEmail, $uniqueEmails)) {
                $uniqueEmails[] = $uniqueEmail;
            }
        }
        return count($uniqueEmails);
    }
}